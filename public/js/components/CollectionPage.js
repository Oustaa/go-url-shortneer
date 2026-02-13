import "./urls/URLItemComponent";

export class CollectionPage extends HTMLElement {
  constructor(endpoint, title) {
    super();
    this.endpoint = endpoint;
    this.title = title;
  }

  async render() {
    const urls = await this.endpoint();
    const ulUrls = this.querySelector("ul");
    ulUrls.innerHTML = "";

    if (urls && urls.length > 0) {
      urls.forEach((url) => {
        const li = document.createElement("li");
        li.appendChild(new URLItemComponent(url));
        ulMovies.appendChild(li);
      });
    } else {
      ulMovies.innerHTML = "<h3>There are no urls</h3>";
    }
  }

  connectedCallback() {
    const template = document.getElementById("template-collection");
    const content = template.content.cloneNode(true);
    this.appendChild(content);

    this.render();
  }
}
