// Styles
import './index.css'
// Fonts
import '@fontsource/roboto/300.css'
import '@fontsource/roboto/400.css'
import '@fontsource/roboto/500.css'
import '@fontsource/roboto/700.css'

import React, { lazy, Suspense } from 'react'
import ReactDOM from 'react-dom/client'
import { createBrowserRouter, Navigate, RouterProvider } from 'react-router-dom'
import { ThemeWrapper } from './theme'
import { client } from './lib/api/apollo'
import { ApolloProvider } from '@apollo/client'
import YamlWorker from './yaml.worker.js?worker'
import { SnackbarProvider } from 'notistack'
import { ReactFlowProvider } from 'reactflow'
import { Box, CircularProgress } from '@mui/material'
import Logo from './components/logo'
import Users from './routes/users'

// Pages
const Root = lazy(() => import('./routes/root'))
const ErrorPage = lazy(() => import('./routes/error-page'))
const Login = lazy(() => import('./routes/auth/login'))
const Blueprints = lazy(() => import('./routes/blueprints'))
const RequestBlueprint = lazy(() => import('./routes/blueprints/request'))
const BlueprintForm = lazy(() => import('./routes/blueprints/form'))
const Deployments = lazy(() => import('./routes/deployments'))
const DestroyDeployment = lazy(() => import('./routes/deployments/destroy'))
const DeploymentDetails = lazy(() => import('./routes/deployments/details'))
const Providers = lazy(() => import('./routes/providers'))
const ProviderForm = lazy(() => import('./routes/providers/form'))
const Permissions = lazy(() => import('./routes/permissions'))
const Projects = lazy(() => import('./routes/projects'))

window.MonacoEnvironment = {
  getWorker(_, label) {
    switch (label) {
      // Handle other cases
      case 'yaml':
        return new YamlWorker()
      default:
        throw new Error(`Unknown label ${label}`)
    }
  },
}

const LazyComponent = ({
  element,
}: {
  element: React.ReactNode
}): React.ReactElement => {
  return (
    <Suspense
      fallback={
        <Box
          sx={{
            width: '100dvw',
            height: '100dvh',
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            justifyContent: 'center',
          }}
        >
          <Box sx={{ position: 'relative' }}>
            <CircularProgress size="4em" />
            <Logo
              sx={{
                width: '1.5em',
                height: '1.5em',
                display: { xs: 'none', md: 'flex' },
                fill: '#F76902',
                position: 'absolute',
                transform: 'translate(-0.75em, -0.8em)',
                top: '50%',
                bottom: '50%',
                left: '50%',
                right: '50%',
              }}
            />
          </Box>
        </Box>
      }
    >
      {element}
    </Suspense>
  )
}

const router = createBrowserRouter([
  {
    path: '/',
    element: <LazyComponent element={<Root />} />,
    children: [
      { index: true, element: <Navigate to="/blueprints" replace /> },
      {
        path: 'blueprints',
        children: [
          { index: true, element: <LazyComponent element={<Blueprints />} /> },
          {
            path: 'create',
            element: (
              <LazyComponent element={<BlueprintForm action="create" />} />
            ),
          },
          {
            path: 'edit/:id',
            element: (
              <LazyComponent element={<BlueprintForm action="edit" />} />
            ),
          },
          {
            path: 'request/:id',
            element: <LazyComponent element={<RequestBlueprint />} />,
          },
        ],
      },
      {
        path: 'deployments',
        children: [
          { index: true, element: <LazyComponent element={<Deployments />} /> },
          {
            path: ':id',
            element: <LazyComponent element={<DeploymentDetails />} />,
          },
          {
            path: 'destroy/:id',
            element: <LazyComponent element={<DestroyDeployment />} />,
          },
        ],
      },
      {
        path: 'providers',
        children: [
          { index: true, element: <LazyComponent element={<Providers />} /> },
          {
            path: 'create',
            element: (
              <LazyComponent element={<ProviderForm action="create" />} />
            ),
          },
          {
            path: 'edit/:id',
            element: <LazyComponent element={<ProviderForm action="edit" />} />,
          },
        ],
      },
      {
        path: 'permissions',
        children: [
          { index: true, element: <LazyComponent element={<Permissions />} /> },
          {
            path: 'create',
            element: (
              <LazyComponent element={<ProviderForm action="create" />} />
            ),
          },
          {
            path: 'edit/:id',
            element: <LazyComponent element={<ProviderForm action="edit" />} />,
          },
        ],
      },
      {
        path: 'projects',
        children: [
          { index: true, element: <LazyComponent element={<Projects />} /> },
        ],
      },
      {
        path: 'users',
        children: [
          { index: true, element: <LazyComponent element={<Users />} /> },
        ],
      },
    ],
    errorElement: <LazyComponent element={<ErrorPage />} />,
  },
  {
    path: '/auth',
    children: [
      {
        path: 'login',
        element: <LazyComponent element={<Login />} />,
      },
    ],
    errorElement: <LazyComponent element={<ErrorPage />} />,
  },
])

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <ThemeWrapper>
        <SnackbarProvider maxSnack={5}>
          <ReactFlowProvider>
            <RouterProvider router={router} />
          </ReactFlowProvider>
        </SnackbarProvider>
      </ThemeWrapper>
    </ApolloProvider>
  </React.StrictMode>
)
