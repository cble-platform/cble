import { Typography, TypographyProps, withStyles } from "@mui/material";

export function TypographyCode(props: TypographyProps) {
  return (
    <Typography
      {...props}
      sx={{
        ...props.sx,
        fontFamily: 'Consolas, Menlo, Monaco, "Andale Mono", "Ubuntu Mono", monospace',
        fontWeight: 400,
        px: "0.5rem",
        py: "0.15rem",
        backgroundColor: "rgba(102, 178, 255, 0.15)",
        borderRadius: 1,
      }}
      component="code"
    >
      {props.children}
    </Typography>
  );
}

// export const TypographySuccess = withStyles(({ theme }) => ({
//   color: theme.palette.error.main,
//   fontWeight: "bold",
// }))(Typography);
