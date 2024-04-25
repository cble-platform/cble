import { Container, Typography } from '@mui/material'
import { useRouteError } from 'react-router-dom'
import Navbar from '../components/navbar'
import { ThemeContext } from '../theme'
import { useContext } from 'react'

export default function ErrorPage() {
  const { themePreference, setThemePreference } = useContext(ThemeContext)
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
      <Navbar themePreference={themePreference} setTheme={setThemePreference} />
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
