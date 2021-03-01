import {precacheAndRoute, getCacheKeyForURL} from "workbox-precaching"
import {registerRoute} from "workbox-routing"
import {CacheFirst, StaleWhileRevalidate} from "workbox-strategies"
import {ExpirationPlugin} from "workbox-expiration"

declare const self: ServiceWorkerGlobalScope

precacheAndRoute(self.__WB_MANIFEST)

registerRoute(
  ({request}) => request.mode === "navigate",
  async () => {
    const index = "/index.html"

    return caches
      .match(getCacheKeyForURL(index) as string)
      .then((response) => {
        return response || fetch(index)
      })
      .catch((err) => {
        console.error(err)

        return fetch(index)
      })
  },
)

registerRoute(
  ({request}) => request.destination === "script" || request.destination === "style",
  new StaleWhileRevalidate({
    cacheName: "code",
  }),
)

registerRoute(
  ({request}) => request.destination === "image",
  new CacheFirst({
    cacheName: "images",
    plugins: [
      new ExpirationPlugin({
        maxEntries: 60,
        maxAgeSeconds: 30 * 24 * 60 * 60, // 30 Days
      }),
    ],
  }),
)
