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
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { ThemeWrapper } from "./theme";
import { client } from "./api/apollo";
import { ApolloProvider } from "@apollo/client";
import EditorTest from "./routes/editor-test";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    children: [
      {
        path: "editor-test",
        element: <EditorTest />,
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
        <RouterProvider router={router} />
      </ThemeWrapper>
    </ApolloProvider>
  </React.StrictMode>
);
