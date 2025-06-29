// Modern JS version (ES6+)
const containers = [];
const lastSeen = {};
const colors = {};
const refreshInterval = Number(document.querySelector('meta[name="refresh-interval"]')?.content || 1000);

function getColor(cowColor) {
  const allColors = [
    "red", "orange", "yellow", "olive", "green", "teal", "blue", "violet", "purple", "pink"
  ];
  let color;
  do {
    color = allColors[Math.floor(Math.random() * allColors.length)];
  } while (color === cowColor);
  return color;
}

async function reload() {
  const ts = Date.now();
  const expireInterval = Number(document.querySelector('meta[name="expire-interval"]')?.content || 10);
  const removeInterval = Number(document.querySelector('meta[name="remove-interval"]')?.content || 20);

  try {
    const response = await fetch(`/demo?${ts}`);
    if (!response.ok) throw new Error('Network response was not ok');
    const data = await response.json();

    if (!(data.version in colors)) {
      colors[data.version] = getColor(data.cowColor);
    }
    const color = colors[data.version];
    if (!containers.includes(data.instance)) {
      containers.push(data.instance);
    }
    const cowColor = data.cowColor;
    lastSeen[data.instance] = ts;

    for (let i = 0; i < containers.length; i++) {
      const instanceName = containers[i];
      let el = document.getElementById(`instance-${instanceName}`);
      if (!el) {
        const match = instanceName.match(/.+-(\w+)$/);
        const displayName = match ? match[1] : instanceName;
        const elData = `<div id="instance-${instanceName}" class="card container-instance"><div class="image" id="replica" style="background-color: ${cowColor};"><img width="25%" height="25%" src="static/img/fish-blue.png"></div><div class="content centered view-computer"><div>${window.showVersion ? `<div class=\"ui top left attached ${color} label\">${data.version}</div>` : ''} <div id="instance-label-${instanceName}"class="ui top right attached yellow serving label">serving</div>${displayName}</div></div></div>`;
        document.querySelector("div.container-group").insertAdjacentHTML('beforeend', elData);
        el = document.getElementById(`instance-${instanceName}`);
      }

      if (ts - lastSeen[instanceName] > refreshInterval) {
        // expire old instances
        const opacity = (1 - (ts - lastSeen[instanceName]) / 1000 / expireInterval).toFixed(2);
        if ((ts - lastSeen[instanceName]) / 1000 > expireInterval + removeInterval) {
          el.remove();
          containers.splice(i, 1);
          delete lastSeen[instanceName];
        } else if (opacity >= 0.2) {
          el.style.opacity = opacity;
        }
      } else {
        el.style.opacity = '';
      }

      // show which replica is active
      const lbl = document.getElementById(`instance-label-${instanceName}`);
      if (instanceName === data.instance) {
        if (lbl) lbl.style.display = '';
      } else {
        if (lbl) lbl.style.display = 'none';
      }
    }

    const countEl = document.getElementById("container-count");
    if (countEl) countEl.textContent = containers.length;
    const labelEl = document.getElementById("container-count-label");
    if (labelEl) labelEl.textContent = containers.length > 1 ? "replicas" : "replica";

    const currEl = document.getElementById("current-container");
    if (currEl) currEl.textContent = data.instance;
    const currVer = document.getElementById("current-container-version");
    if (currVer) currVer.textContent = data.version;
    const backend = document.getElementById("container-backend");
    if (backend) backend.classList.remove('hidden');
    document.querySelectorAll("div.container-backend").forEach(div => div.style.display = '');
    const extraInfo = document.getElementById("extra-info");
    if (extraInfo) extraInfo.textContent = data.metadata;

    const reqCount = document.getElementById("requests-count");
    if (reqCount) {
      const current = parseInt(reqCount.textContent, 10);
      reqCount.textContent = (isNaN(current) ? 0 : current) + 1;
    }
    // If errors is 0, keep it green, else red
    const errCountZero = document.getElementById("requests-error-count");
    if (errCountZero) {
      if (errCountZero.textContent === "0") {
        errCountZero.classList.remove("text-red-500");
        errCountZero.classList.add("text-green-500");
      } else {
        errCountZero.classList.remove("text-green-500");
        errCountZero.classList.add("text-red-500");
      }
    }
  } catch (err) {
    const errCount = document.getElementById("requests-error-count");
    if (errCount) {
      const current = parseInt(errCount.textContent, 10);
      const newVal = (isNaN(current) ? 0 : current) + 1;
      errCount.textContent = newVal;
      if (newVal === 0) {
        errCount.classList.remove("text-red-500");
        errCount.classList.add("text-green-500");
      } else {
        errCount.classList.remove("text-green-500");
        errCount.classList.add("text-red-500");
      }
    }
    const errStat = document.querySelector(".error");
    if (errStat) errStat.classList.add("red");
  }
}

setInterval(reload, refreshInterval);
