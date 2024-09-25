import './projector/scale.js';

const eventSource = new EventSource(`/system/projector/subscribe/${window.location.href.substring(window.location.href.lastIndexOf('/') + 1)}`)
eventSource.addEventListener(`settings`, (e) => {
  console.log(e, JSON.parse(e.data));
});

eventSource.addEventListener(`deleted`, (e) => {
  console.log(e);
});

window.addEventListener(`unload`, () => {
  eventSource.close();
});
