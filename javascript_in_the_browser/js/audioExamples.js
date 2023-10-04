const AudioContext = window.AudioContext || window.webKitAudioContext
const audioCtx = new AudioContext();

const audioElement = document.querySelector("audio");
const playBtn = document.querySelector("button");
const volumeSlider = document.querySelector(".volume");

const audioSource = audioCtx.createMediaStreamSource(audioElement);

// Next we include a couple of event handlers that serve to toggle
// between play and pause when the button is pressed and reset
// the display back to the begining when the song has finished
// playing
playBtn.addEventListener("click", ()=>{
  // check if context is in suspended state (autoplay policy)
  if(audioCtx.state === "suspended"){
    audioCtx.resume();
  }

  // if track is stopped, play it
  if(playBtn.getAttribute("class") === "paused"){
    audioElement.play();
    playBtn.setAttribute("class", "playing");
    playBtn.textContent = "Pause";
  } else if (playBtn.getAttribute("class") === "playing"){
    audioElement.pause();
    playBtn.setAttribute("class", "paused");
    playBtn.textContent = "Play"
  }
});

// If track ends
audioElement.addEventListener("ended", () =>{
  playBtn.setAttribute("class", "paused");
  playBtn.textContent = "play";
})

// Next, we create a GainMode object using AudioContext.createGain()
// method, which can be used to adjust the volume of audio fed
// through it, and create another event handler that changes the
// value of the audio graph's gain volume

const gainMode = audioCtx.createGain();

volumeSlider.addEventListener("input", ()=>{
  gainMode.gain.value= volumeSlider.value;
});

// The final thing to do to get this to work is to connect
// the different nodes in the audio graph up, which is done
// using the AudioNode.connect() method available on
// every node type
audioSource.connect(gainMode).connect(audioCtx.destination);
