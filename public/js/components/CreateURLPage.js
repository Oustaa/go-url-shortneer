import { API } from "../../services/api";
import { CollectionPage } from "../CollectionPage";

export class UrlsPage extends CollectionPage {
  constructor() {
    super(API.getUrls, "Urls List");
  }
}
