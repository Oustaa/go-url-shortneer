import { API } from "../services/API.js";
import { CollectionPage } from "../components/CollectionPage.js";

export class URLsPage extends CollectionPage {
  constructor() {
    super(API.getUrls, "Urls List");
  }

  connectedCallback() {
    const template = document.getElementById("template-URLs");
    const content = template.content.cloneNode(true);

    this.appendChild(content);
  }
}

customElements.define("urls-page", URLsPage);
