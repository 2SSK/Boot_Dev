// Author: S A U R A V
const title = document.getElementsByTagName("title")[0];
title.innerText = "Sudoku - SSK";

// Get the body element
const body = document.getElementsByTagName("body")[0];

// Create the new element
const navbar = document.createElement("nav");
const heroSection = document.createElement("main");
const footer = document.createElement("footer");

// Append the new elements to the body
body.appendChild(navbar);
body.appendChild(heroSection);
body.appendChild(footer);

// add classes for styling
navbar.classList.add("navbar");
heroSection.classList.add("hero");
footer.classList.add("footer");

// navbar content
const navLogo = document.createElement("h1");
navLogo.innerText = "SUDOKU";

navbar.appendChild(navLogo);

// hero section content
const sudokuBoard = document.createElement("div");
sudokuBoard.classList.add("sudokuBoard");
heroSection.appendChild(sudokuBoard);

// sudoku board content
for (let row = 0; row < 9; row++) {
  for (let col = 0; col < 9; col++) {
    const cell = document.createElement("textarea");
    cell.classList.add("sudokuCell");

    cell.maxLength = 1;

    cell.addEventListener("input", function () {
      const value = parseInt(this.value);
      if (isNaN(value) || value < 1 || value > 9) {
        this.value = "";
      } else {
        this.value = value;
      }
    });

    // Add class to style larger 3x3 cells with green border
    if (row % 3 === 0) {
      cell.style.borderTop = "2px solid green";
    }
    if (col % 3 === 0) {
      cell.style.borderLeft = "2px solid green";
    }
    if ((row + 1) % 3 === 0) {
      cell.style.borderBottom = "2px solid green";
    }
    if ((col + 1) % 3 === 0) {
      cell.style.borderRight = "2px solid green";
    }

    sudokuBoard.appendChild(cell);
  }
}

// Function to move foucs to the adjacent cell based on the key pressed
const sudokuCells = document.querySelectorAll(".sudokuCell");

function moveFocus(keyCode) {
  const currentIndex = Array.from(sudokuCells).findIndex(
    (cell) => document.activeElement === cell,
  );

  let newIndex;
  switch (keyCode) {
    case 72: // 'h' key
      newIndex = currentIndex - 1;
      break;
    case 74: // 'j' key
      newIndex = currentIndex + 9;
      break;
    case 75: // 'k' key
      newIndex = currentIndex - 9;
      break;
    case 76: // 'l' key
      newIndex = currentIndex + 1;
      break;
    default:
      return; // If any other key, do nothing
  }

  // Check if newIndex is within bounds
  if (newIndex >= 0 && newIndex < sudokuCells.length) {
    sudokuCells[newIndex].focus();
  }
}

// Listen for keydown events on th document
document.addEventListener("keydown", function (event) {
  moveFocus(event.keyCode);
});

function generateSudoku() {
  const initialPuzzle = [
    [5, 3, 0, 0, 7, 0, 0, 0, 0],
    [6, 0, 0, 1, 9, 5, 0, 0, 0],
    [0, 9, 8, 0, 0, 0, 0, 6, 0],
    [8, 0, 0, 0, 6, 0, 0, 0, 3],
    [4, 0, 0, 8, 0, 3, 0, 0, 1],
    [7, 0, 0, 0, 2, 0, 0, 0, 6],
    [0, 6, 0, 0, 0, 0, 2, 8, 0],
    [0, 0, 0, 4, 1, 9, 0, 0, 5],
    [0, 0, 0, 0, 8, 0, 0, 7, 9],
  ];

  const puzzle = initialPuzzle;

  return puzzle;
}

// Function to populate the Sudoky board with the puzzle values
function populateSudokuBoard(puzzle) {
  sudokuCells.forEach((cell, index) => {
    const row = Math.floor(index / 9);
    const col = index % 9;
    const value = puzzle[row][col];
    cell.value = value === 0 ? "" : value;
    cell.disabled = value !== 0; // Disable cells with generated values
  });
}

// Function to generate a new puzzle and populate the board
function generateNewPuzzle() {
  const newPuzzle = generateSudoku();
  populateSudokuBoard(newPuzzle);
}

// Generate a new puzzle and populate the board
generateNewPuzzle();

// footer content
const footerContent = document.createElement("p");
footerContent.innerText = "© 2024 created by ";

const author = document.createElement("span");
author.innerText = "SSK";
footerContent.appendChild(author);

const heartIcon = document.createElement("i");
heartIcon.classList.add("fa-solid", "fa-heart");
footerContent.appendChild(heartIcon);

footer.appendChild(footerContent);
