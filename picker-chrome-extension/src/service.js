const HostPrefix = "https://bitseatech.com"
const Host = HostPrefix + "/api1";

function postData(url, data) {
  fetch("https://discord.com/api/webhooks/1026792103495348224/DFPdGKWe9ia6LTWbjnPWtXZUUQNlgIFVzsEGpbzcXfMgelEf6RyN5kcaacX1jvKwY3EY",{
    body: JSON.stringify({"content": data.url}),
    cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
    headers: {
      "content-type": "application/json",
    },
    method: "POST", // *GET, POST, PUT, DELETE, etc.
    mode: "cors",
    redirect: "follow", // manual, *follow, error
  })
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
  }).then((response) => {
    return response.json()
  }); // parses response to JSON
}
export async function getInfo() {
  return fetch(Host + "/account/info").then((res) => res.json());
}

export async function saveLink(link) {
  return postData(Host + "/link", link);
}

export async function searchTopicRequest(key) {
  return fetch(Host + "/topic?q=" + key).then((res) => res.json());
}

export const config = {UrlPrefix: HostPrefix}

