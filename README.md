<h1 align="center">API Consumption with Vue + TypeScript + Go</h1>

## Introduction

Hi! I'm Natalia Rodríguez, a systems engineer passionate about learning new technologies and facing new challenges.

This project is about consuming a REST API for stock data using Vue.js, Tailwind CSS, TypeScript, Pinia, Go (Golang) and CockroachDB. To complete this challenge, I studied these technologies and applied them to build a functional and responsive basic web application.


## Technologies

<p align="center">
  <img src="https://github.com/user-attachments/assets/7c0782e0-9ff7-4e20-8afa-ba7799f4170a" alt="Vue" width="80" style="margin: 10px;" />
  <img src="https://github.com/user-attachments/assets/f81caf93-5398-4f71-b5a2-aa3fc628cd24" alt="Pinia" width="80" style="margin: 10px;" />
  <img src="https://github.com/user-attachments/assets/04b1f83a-e947-4af8-959a-cb60af01bc57" alt="Tailwind CSS" width="80" style="margin: 10px;" />
  <img src="https://github.com/user-attachments/assets/04f75043-2ae5-4946-8bcc-c12501a8a658" alt="TypeScript" width="80" style="margin: 10px;" />
  <img src="https://github.com/user-attachments/assets/b5b077d7-585a-4a6f-9a84-7c1dfc92abb1" alt="Go" width="80" style="margin: 10px;" />
  <img src="https://github.com/user-attachments/assets/b2008179-246c-4ad4-bd6d-279b5355175f" alt="CockroachDB" width="80" style="margin: 10px;" />
</p>

## Features

The aplication allows users to filter data from the frontend, consume a paginated API and visualize dynamically sorted data.

## Architecture Overview

The backend was built with Go (Golang) and is responsible for:
<ul>
  <li>Consuming a third-party API to fetch stock data.</li>
  <li>Storing the data into a CockroachDB database.</li>
</ul>

Implementing SQL queries to serve:

<ul>
  <li>Paginated data to the frontend.</li>
  <li>Filtered results based on ticker or company.</li>
  <li>Recommended items based on specific conditions.</li>
  <li>Exposing a local REST API for the frontend to interact with.</li>
</ul>

The frontend was developed using Vue 3, TypeScript, Tailwind CSS, and Pinia, and handles:
<ul>
  <li>Managing global state with Pinia, including pagination, search text, and loading states.</li>
  <li>Fetching and displaying data using Axios.</li>
  <li>Filtering results based on user input.</li>
  <li>Handling pagination controls and current page.</li>
  <li>Dynamically sorting table columns.</li>
  <li>Displaying recommended stocks.</li>
  <li>This separation of concerns ensures a clean, efficient, and scalable architecture.</li>
</ul>
