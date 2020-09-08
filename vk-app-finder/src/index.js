import "core-js/features/map";
import "core-js/features/set";
import React from "react";
import ReactDOM from "react-dom";
import bridge from "@vkontakte/vk-bridge";
import App from "./App";
import { RouterProvider } from 'react-router5';
import createRouter from './create-router';

bridge.send("VKWebAppInit");
const search = window.location.search;
const router = createRouter().start();

ReactDOM.render((
  <RouterProvider router={router}>
    <App router={router} search={search}/>
  </RouterProvider>
  ), document.getElementById("root")
);

if (process.env.NODE_ENV === "development") {
  import("./eruda").then(({ default: eruda }) => {}); //runtime download
}
