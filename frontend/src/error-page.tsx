import { Container, Typography } from '@mui/material'
import { useRouteError } from 'react-router-dom'

export default function ErrorPage() {
  const error = useRouteError() as {
    statusText?: string
    error?: Error
  }
  console.error(error?.error ?? error)

  return (
    <Container
      sx={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        flexDirection: 'column',
        minHeight: '100dvh',
      }}
    >
      <Typography variant="h3">Oops!</Typography>
      <Typography variant="body1">
        Sorry, an unexpected error has occurred.
      </Typography>
      <Typography variant="subtitle2" color="error">
        {<i>{error?.error?.message || error.statusText}</i>}
      </Typography>
    </Container>
  )
}
