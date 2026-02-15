import { CreateURLPage } from "../components/CreateURLPage.js";
import { LoginPage } from "../components/LoginPage.js";
import { RegisterPage } from "../components/RegisterPage.js";
import { URLsPage } from "../components/URLsPage.js";

export const routes = [
  {
    path: "/",
    component: URLsPage,
    loggedIn: true,
  },
  {
    path: "/urls/create",
    component: CreateURLPage,
  },
  {
    path: "/account/login",
    component: LoginPage,
  },
  {
    path: "/account/register",
    component: RegisterPage,
  },
];
