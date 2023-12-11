import { AppBar, Box, Button, Container, IconButton, Toolbar, Typography } from "@mui/material";
import Logo from "./logo";
import { Brightness4, Brightness7, BrightnessAuto } from "@mui/icons-material";

export default function Navbar({
  themePreference,
  setTheme,
}: {
  themePreference: "light" | "dark" | "auto";
  setTheme: (theme: "auto" | "dark" | "light") => void;
}) {
  return (
    <AppBar position="fixed" color="primary">
      <Container maxWidth="xl">
        <Toolbar disableGutters>
          <Logo sx={{ display: { xs: "none", md: "flex" }, mr: 1, fill: "white" }} />
          <Typography
            variant="h6"
            noWrap
            component="div"
            sx={{
              mr: 2,
              display: { xs: "none", md: "flex" },
              fontWeight: 700,
              color: "white",
              textDecoration: "none",
            }}
          >
            CBLE
          </Typography>
          {/* <Box sx={{ flexGrow: 1, display: { xs: "flex", md: "none" } }}>
            <IconButton
              size="large"
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              onClick={handleOpenNavMenu}
              color="inherit"
            >
              <MenuIcon />
            </IconButton>
            <Menu
              id="menu-appbar"
              anchorEl={anchorElNav}
              anchorOrigin={{
                vertical: "bottom",
                horizontal: "left",
              }}
              keepMounted
              transformOrigin={{
                vertical: "top",
                horizontal: "left",
              }}
              open={Boolean(anchorElNav)}
              onClose={handleCloseNavMenu}
              sx={{
                display: { xs: "block", md: "none" },
              }}
            ></Menu>
          </Box>
          <AdbIcon sx={{ display: { xs: 'flex', md: 'none' }, mr: 1 }} /> */}
          <Box sx={{ flexGrow: 1, display: { xs: "none", md: "flex" } }}>
            <Button href="/blueprints" sx={{ my: 2, color: "white", display: "block", textAlign: "center" }}>
              Blueprints
            </Button>
            <Button href="/deployments" sx={{ my: 2, color: "white", display: "block", textAlign: "center" }}>
              Deployments
            </Button>
            <Button href="/providers" sx={{ my: 2, color: "white", display: "block", textAlign: "center" }}>
              Providers
            </Button>
          </Box>
          <IconButton
            color="inherit"
            onClick={() => {
              if (themePreference === "light") setTheme("dark");
              else if (themePreference === "dark") setTheme("auto");
              else setTheme("light");
            }}
          >
            {themePreference === "light" && <Brightness4 />}
            {themePreference === "dark" && <Brightness7 />}
            {themePreference === "auto" && <BrightnessAuto />}
          </IconButton>
        </Toolbar>
      </Container>
    </AppBar>
  );
}
