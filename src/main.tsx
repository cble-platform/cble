// Styles
import "./index.css";
// Fonts
import "@fontsource/roboto/300.css";
import "@fontsource/roboto/400.css";
import "@fontsource/roboto/500.css";
import "@fontsource/roboto/700.css";

// Pages
import Root from "./routes/root";
import ErrorPage from "./error-page";
import Login from "./routes/auth/login";

import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, Navigate, RouterProvider } from "react-router-dom";
import { ThemeWrapper } from "./theme";
import { client } from "./api/apollo";
import { ApolloProvider } from "@apollo/client";
import Blueprints from "./routes/blueprints";
import YamlWorker from "./yaml.worker.js?worker";
import { SnackbarProvider } from "notistack";
import RequestBlueprint from "./routes/blueprints/request";
import BlueprintForm from "./routes/blueprints/form";
import Deployments from "./routes/deployments";
import DestroyDeployment from "./routes/deployments/destroy";

window.MonacoEnvironment = {
  getWorker(moduleId, label) {
    switch (label) {
      // Handle other cases
      case "yaml":
        return new YamlWorker();
      default:
        throw new Error(`Unknown label ${label}`);
    }
  },
};

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    children: [
      { index: true, element: <Navigate to="/blueprints" replace /> },
      {
        path: "blueprints",
        children: [
          { index: true, element: <Blueprints /> },
          { path: "create", element: <BlueprintForm action="create" /> },
          { path: "edit/:id", element: <BlueprintForm action="edit" /> },
          { path: "request/:id", element: <RequestBlueprint /> },
        ],
      },
      {
        path: "deployments",
        children: [
          { index: true, element: <Deployments /> },
          { path: "destroy/:id", element: <DestroyDeployment /> },
        ],
      },
    ],
    errorElement: <ErrorPage />,
  },
  {
    path: "/auth",
    children: [
      {
        path: "login",
        element: <Login />,
      },
    ],
    errorElement: <ErrorPage />,
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <ThemeWrapper>
        <SnackbarProvider maxSnack={5}>
          <RouterProvider router={router} />
        </SnackbarProvider>
      </ThemeWrapper>
    </ApolloProvider>
  </React.StrictMode>
);
