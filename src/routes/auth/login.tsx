import { LockOutlined } from "@mui/icons-material";
import { Avatar, Box, IconButton, Checkbox, Container, FormControlLabel, Grid, Link, TextField, Typography } from "@mui/material";
import { useContext, useState } from "react";
import { Login as ApiLogin } from "../../api/auth";
import { useNavigate } from "react-router-dom";
import { LoadingButton } from "@mui/lab";
import { ThemeContext } from "../../theme";
import { Brightness4, Brightness7, BrightnessAuto } from "@mui/icons-material";

export default function Login() {
  const [loginData, setLoginData] = useState<{
    username: string;
    password: string;
  }>({ username: "", password: "" });
  const [loginLoading, setLoginLoading] = useState<boolean>(false);
  const navigate = useNavigate();
  const { themePreference, setThemePreference } = useContext(ThemeContext);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (loginData.username == null || loginData.password == null) return;

    setLoginLoading(true);
    ApiLogin(loginData.username, loginData.password)
      .then((res) => {
        if (res.status !== 200) {
          console.error(`Auth failed with status ${res.statusText}`);
        } else {
          navigate("/");
        }
      })
      .catch(console.error)
      .finally(() => setLoginLoading(false));
  };

  return (
    <Container component="main" maxWidth="xs">
      <IconButton
        color="inherit"
        onClick={() => {
          if (themePreference === "light") setThemePreference("dark");
          else if (themePreference === "dark") setThemePreference("auto");
          else setThemePreference("light");
        }}
        sx={{
          position: "absolute",
          top: "1rem",
          right: "1rem",
        }}
      >
        {themePreference === "light" && <Brightness4 />}
        {themePreference === "dark" && <Brightness7 />}
        {themePreference === "auto" && <BrightnessAuto />}
      </IconButton>
      <Box
        sx={{
          // marginTop: 8,
          minHeight: "100dvh",
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
        }}
      >
        <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
          <LockOutlined />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign in
        </Typography>
        <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
          <TextField
            margin="normal"
            required
            fullWidth
            id="username"
            label="Username"
            name="username"
            autoComplete="username"
            autoFocus
            value={loginData.username}
            onChange={(e) => setLoginData((prevVal) => ({ ...prevVal, username: e.target.value }))}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
            value={loginData.password}
            onChange={(e) => setLoginData((prevVal) => ({ ...prevVal, password: e.target.value }))}
          />
          <FormControlLabel control={<Checkbox value="remember" color="primary" />} label="Remember me" />
          <LoadingButton
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
            disabled={loginData.username === "" || loginData.password === ""}
            loading={loginLoading}
            loadingPosition="start"
          >
            Sign In
          </LoadingButton>
          <Grid container>
            <Grid item xs>
              <Link href="#" variant="body2">
                Forgot password?
              </Link>
            </Grid>
            {/* <Grid item>
              <Link href="#" variant="body2">
                {"Don't have an account? Sign Up"}
              </Link>
            </Grid> */}
          </Grid>
        </Box>
      </Box>
    </Container>
  );
}
