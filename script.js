const titles = [
  "Interested?",
  "Like What You See?",
  "Get In Touch!",
  "Reach Out Today!",
  "Let's Connect!",
  "Tap to Contact!",
  "Contact Us Now!",
  "Get Started!",
  "Learn More!",
  "Find Out More!",
  "Click to Get in Touch!",
  "Let's Talk!",
  "Start Today!",
];

linkButton = document.getElementsByClassName("card linked");

if (linkButton && linkButton.length == 1) {
  buttonTitle = linkButton[0].getElementsByTagName("h3")[0];
  buttonTitle.innerText = titles[Math.floor(Math.random() * titles.length)];
}
