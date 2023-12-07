import { Link, Typography } from "@mui/material";

function Copyright(props: any) {
  return (
    <Typography variant="body2" color="text.secondary" align="center" {...props}>
      {"Copyright © "}
      <Link color="inherit" href="https://github.com/cble-platform">
        CBLE Platform
      </Link>{" "}
      {new Date().getFullYear()}
      {"."}
    </Typography>
  );
}
