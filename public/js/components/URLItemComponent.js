export class URLItemComponent extends HTMLElement {
  constructor(url) {
    super();
    this.url = url;
  }

  connectedCallback() {
    this.innerHTML = `<h4>${this.url.long_url}</h4>
  ${this.url.user && `<p>user: ${this.url.user.username}</p>`}`;
  }
}

customElements.define("url-item", URLItemComponent);
