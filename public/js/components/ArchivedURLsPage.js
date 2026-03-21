import { API } from "../services/API.js";

export class ArchivedURLsPage extends HTMLElement {
  constructor() {
    super();
    this.loading = false;
  }

  async connectedCallback() {
    this.loading = true;

    const template = document.getElementById("template-URLs");
    const content = template.content.cloneNode(true);

    const urls = await API.getArchivedUrls();

    console.log({ urls });

    const table = document.createElement("table");

    const thead = document.createElement("thead");
    const headerRow = document.createElement("tr");

    ["Short URL", "Long URL", "Vists", "Actions"].forEach((text) => {
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

      const vistsCell = document.createElement("td");
      vistsCell.textContent = element.visits;

      const actionsCell = document.createElement("td");

      const restoreButton = document.createElement("button");
      restoreButton.textContent = "Restore";

      restoreButton.addEventListener("click", async () => {
        await API.restoreURL(element);
      });

      actionsCell.append(restoreButton);

      row.append(shortCell, longCell, vistsCell, actionsCell);
      tbody.appendChild(row);
    });

    table.appendChild(tbody);

    content.appendChild(table);
    this.appendChild(content);

    this.loading = false;
  }
}

customElements.define("archived-urls-page", ArchivedURLsPage);
