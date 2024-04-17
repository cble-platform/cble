import { Fab, FabProps } from '@mui/material'

export default function ContainerFab(props: FabProps) {
  return (
    <Fab
      {...props}
      sx={{
        ...props.sx,
        position: 'fixed',
        right: {
          xs: '16px',
          sm: '24px',
          lg: 'calc((100dvw - 1200px) / 2 + 24px)',
        },
        bottom: '2rem',
      }}
    >
      {props.children}
    </Fab>
  )
}
