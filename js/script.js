// Copyright (c) 2023 Jaden Plugowsky All rights reserved
//
// Created by: Jaden Plugowsky
// Created on: April 2023
// This file contains the JS functions for index.html

"use strict"

const randomNumber = Math.floor(Math.random() * 6) + 1

function guessPressed() {
  //This function takes a user's guess and checks if it is equal to the previously randomly-generated number
  //Input through Textfields
  const guessedNumber = parseInt(
    document.getElementById("guessed-number").value
  )

  //Process
  //pass

  //Output
  if (guessedNumber == randomNumber) {
    document.getElementById("answer").innerHTML =
      "Congratulations, you have guessed " +
      guessedNumber +
      ", which is the correct number!"
  }
  if (guessedNumber == guessedNumber.isNaN()) {
    document.getElementById("answer").innerHTML =
      "Try again, that answer is not a number."
  }
  if (guessedNumber != randomNumber) {
    document.getElementById("answer").innerHTML =
      "Try again, " + guessedNumber + " is not the correct number."
  }
}
