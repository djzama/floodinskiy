var speedometer = document.getElementById("speedometer");
var packetCount = 0;
var startTime = Date.now();

setInterval(function() {
    var currentTime = Date.now();
    var elapsedSeconds = (currentTime - startTime) / 1000;
    var packetsPerSecond = (packetCount / elapsedSeconds).toFixed(2);
    speedometer.textContent = packetsPerSecond + " pps";
}, 1000);

document.querySelector("form").addEventListener("submit", function() {
    packetCount += parseInt(document.querySelector("#packetCount").value);
});
