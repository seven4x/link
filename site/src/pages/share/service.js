function postData(url, data) {
  // Default options are marked with *
  return fetch(url, {
    body: JSON.stringify(data), // must match 'Content-Type' header
    cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
    credentials: "same-origin", // include, same-origin, *omit
    headers: {
      "content-type": "application/json",
    },
    method: "POST", // *GET, POST, PUT, DELETE, etc.
    mode: "cors", // no-cors, cors, *same-origin
    redirect: "follow", // manual, *follow, error
  }).then((response) => response.json()); // parses response to JSON
}

export async function getInfo() {
  return fetch("/api1/account/info").then((res) => res.json());
}

export async function saveLink(link) {
  return postData("/api1/link", link);
}

export async function searchTopicRequest(key) {
  return fetch("/api1/topic?q=" + key).then((res) => res.json());
}


