package installer

import (
	"os"
	"path/filepath"

	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/plan"
)

const zustandStorePath = "src/stores/useSparkyStore.ts"

const zustandStoreContent = `import { create } from 'zustand';

type SparkyState = {
  treats: number;
  hype: boolean;
  mood: string;
};

type SparkyActions = {
  addTreat: () => void;
  reset: () => void;
  toggleHype: () => void;
};

type SparkyStore = SparkyState & SparkyActions;

const BASE_MOOD = 'ready to ship';
const HYPE_MOOD = 'buzzing to ship';

export const useSparkyStore = create<SparkyStore>((set) => ({
  treats: 1,
  hype: true,
  mood: HYPE_MOOD,
  addTreat: () =>
    set((state) => ({
      treats: state.treats + 1,
      mood: state.hype ? HYPE_MOOD : BASE_MOOD,
    })),
  reset: () => set({ treats: 1, hype: true, mood: HYPE_MOOD }),
  toggleHype: () =>
    set((state) => {
      const hype = !state.hype;
      return {
        hype,
        mood: hype ? HYPE_MOOD : BASE_MOOD,
      };
    }),
}));
`

// InstallZustand installs Zustand dependency.
func InstallZustand(p plan.Plan) error {
	spin := logger.StartSpinner("Installing Zustand")
	if err := addDependencies(p, false, "zustand@latest"); err != nil {
		spin("Failed to install Zustand")
		return err
	}
	spin("Installed Zustand")
	return nil
}

// RemoveZustand uninstalls Zustand.
func RemoveZustand(p plan.Plan) error {
	spin := logger.StartSpinner("Removing Zustand")
	if err := removeDependencies(p, false, "zustand"); err != nil {
		spin("Failed to remove Zustand")
		return err
	}
	spin("Removed Zustand")
	return nil
}

// WriteZustandStore writes a demo Zustand store used by the default App.
func WriteZustandStore() error {
	_, err := writeZustandStore(false)
	return err
}

// WriteZustandStoreIfMissing writes the demo store when it is absent.
func WriteZustandStoreIfMissing() (bool, error) {
	return writeZustandStore(true)
}

func writeZustandStore(skipIfExists bool) (bool, error) {
	if skipIfExists {
		if _, err := os.Stat(zustandStorePath); err == nil {
			return false, nil
		} else if !os.IsNotExist(err) {
			return false, err
		}
	}

	if err := os.MkdirAll(filepath.Dir(zustandStorePath), 0o755); err != nil {
		return false, err
	}

	if err := os.WriteFile(zustandStorePath, []byte(zustandStoreContent), 0o644); err != nil {
		return false, err
	}

	return true, nil
}

// DeleteZustandStoreIfOwned removes the demo store when it matches the generated content.
func DeleteZustandStoreIfOwned() error {
	data, err := os.ReadFile(zustandStorePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	if string(data) != zustandStoreContent {
		return nil
	}

	if err := os.Remove(zustandStorePath); err != nil {
		return err
	}

	_ = os.Remove(filepath.Dir(zustandStorePath))
	return nil
}
