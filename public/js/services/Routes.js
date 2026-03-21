import { CreateURLPage } from "../components/CreateURLPage.js";
import { LoginPage } from "../components/LoginPage.js";
import { RegisterPage } from "../components/RegisterPage.js";
import { URLsPage } from "../components/URLsPage.js";
import { ArchivedURLsPage } from "../components/ArchivedURLsPage.js";
import { LogoutPage } from "../components/LogoutPage.js";

export const routes = [
  {
    path: "/",
    component: URLsPage,
    loggedIn: true,
  },
  {
    path: "/archived",
    component: ArchivedURLsPage,
    loggedIn: true,
  },
  {
    path: "/urls/create",
    component: CreateURLPage,
  },
  // auth
  {
    path: "/account/login",
    component: LoginPage,
  },
  {
    path: "/account/register",
    component: RegisterPage,
  },
  {
    path: "/account/logout",
    component: LogoutPage,
    loggedIn: true,
  },
];
