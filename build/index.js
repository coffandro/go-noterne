const overlay = document.getElementById('loading-overlay');
function hideOverlay() { overlay.style.display = 'none' }
function showOverlay() { overlay.style.display = 'flex' }

// show loading overlay
showOverlay();

import Module from "./rl/raylib.js";

const canvas = document.getElementById('canvas');


const wasmBinary = await fetch("./rl/raylib.wasm")
  .then(r => r.arrayBuffer());

// disable right click context menu
canvas.addEventListener('contextmenu', e => e.preventDefault())

globalThis.onRuntimeInitialized = function onRuntimeInitialized() {
  console.log("HI!");
  function resizeCanvasAndClient() {
    var width = window.innerWidth | 0;
    var height = window.innerHeight | 0;

    canvas.width = width;
    canvas.height = height;
    canvas.style.width = width + 'px';
    canvas.style.height = height + 'px';

    console.log(width, height);

    if (typeof goSetDimensions === 'function') {
      goSetDimensions(width, height);
    }
  }


  window.addEventListener('resize', resizeCanvasAndClient);
  resizeCanvasAndClient();
}

let mod = await Module({
  canvas,
  wasmBinary: new Uint8Array(wasmBinary),
  print: (text) => {
    console.log(text);
  }
});

window.mod = mod;
const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
  .then(result => {
    // hide loading overlay before running code
    hideOverlay();
    go.run(result.instance);
  })
  .catch(console.error);
