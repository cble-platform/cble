import { Container } from "@mui/material";
import { useContext, useEffect } from "react";
import { Outlet, useNavigate } from "react-router-dom";
import Navbar from "../components/navbar";
import { ThemeContext } from "../theme";
import { useMeQuery } from "../api/graphql/generated";

export default function Root() {
  const { themePreference, setThemePreference } = useContext(ThemeContext);
  const { data: meData, loading: meLoading, error: meError } = useMeQuery();
  const navigate = useNavigate();

  useEffect(() => {
    if (meError && meError?.networkError?.message === "Response not successful: Received status code 401") navigate("/auth/login");
  }, [meData, meLoading, meError]);

  return (
    <>
      <Navbar themePreference={themePreference} setTheme={setThemePreference} />
      <Container>
        <Outlet />
      </Container>
    </>
  );
}
