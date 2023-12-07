import { useMediaQuery, createTheme, ThemeProvider, CssBaseline } from "@mui/material";
import { useState, useMemo, createContext, useEffect } from "react";
import { CBLEPalette } from "./colors";

export const ThemeContext = createContext<{
  themePreference: "light" | "dark" | "auto";
  setThemePreference: (theme: "light" | "dark" | "auto") => void;
}>({ themePreference: "auto", setThemePreference: () => undefined });

export function ThemeWrapper({ children }: { children: React.ReactElement }) {
  const [themePreference, setThemePreference] = useState<"light" | "dark" | "auto">(
    (localStorage.getItem("theme") as "light" | "dark" | "auto") ?? "auto"
  );
  const prefersDarkMode = useMediaQuery("(prefers-color-scheme: dark)");

  const theme = useMemo(
    () =>
      createTheme({
        palette: {
          mode: themePreference === "auto" ? (prefersDarkMode ? "dark" : "light") : themePreference,
          ...CBLEPalette,
        },
      }),
    [prefersDarkMode, themePreference]
  );

  return (
    <ThemeContext.Provider
      value={{
        themePreference: themePreference,
        setThemePreference: (theme) => {
          localStorage.setItem("theme", theme);
          setThemePreference(theme);
        },
      }}
    >
      <ThemeProvider theme={theme}>
        <CssBaseline />
        {children}
      </ThemeProvider>
    </ThemeContext.Provider>
  );
}
