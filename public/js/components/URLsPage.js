import { API } from "../services/API.js";

export class URLsPage extends HTMLElement {
  constructor() {
    super();
    this.loading = false;
  }

  async connectedCallback() {
    this.loading = true;

    const template = document.getElementById("template-URLs");
    const content = template.content.cloneNode(true);

    const urls = await API.getUrls();

    const table = document.createElement("table");

    const thead = document.createElement("thead");
    const headerRow = document.createElement("tr");

    ["Short URL", "Long URL"].forEach((text) => {
      const th = document.createElement("th");
      th.textContent = text;
      headerRow.appendChild(th);
    });

    thead.appendChild(headerRow);
    table.appendChild(thead);

    const tbody = document.createElement("tbody");

    urls.forEach((element) => {
      const row = document.createElement("tr");

      const shortCell = document.createElement("td");
      shortCell.textContent = element.short_url;

      const longCell = document.createElement("td");
      longCell.textContent = element.long_url;

      row.append(shortCell, longCell);
      tbody.appendChild(row);
    });

    table.appendChild(tbody);

    content.appendChild(table);
    this.appendChild(content);

    this.loading = false;
  }
}

customElements.define("urls-page", URLsPage);
