<div id="top"></div>



<!-- PROJECT LOGO -->
<br />
<div align="center">
    <img src="./ressources/logo.gif" alt="Logo" width="80" height="80">

<h3 align="center">A*</h3>

  <p align="center">Learn go concurrence with goroutines using design pattern (Semaphore, pipeline ...)</p>
</div>

 <br />  

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">🧭 About The Project</a>
      <ul>
        <li><a href="#built-with">🏗️ Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">📋 Getting Started</a>
      <ul>
        <li><a href="#prerequisites">🗺️ Prerequisites</a></li>
        <li><a href="#installation">⚙️ Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">💾 Usage</a></li>
    <li><a href="#contributing">🔗 Contributing</a></li>
    <li><a href="#license">📰 License</a></li>
    <li><a href="#contact">📫 Contact</a></li>
    <li><a href="#acknowledgments">⛱️ Acknowledgments</a></li>
  </ol>
</details>

<br>


<!-- ABOUT THE PROJECT -->
## 🧭 About The Project

### 🏗️ Built With

List of technologies / frameworks used during the project.

* [![Go Badge](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
* [Concurrence](https://go.dev/tour/concurrency/1)
<p align="right"><a href="#top">⬆️</a></p>

<!-- GETTING STARTED -->
## 📋 Getting Started

This project is a implementation of A* algorithm in Go. The goal is to compare the performance of the algorithm in Go and C++.
Implementation in C++ is not asynchronous unlike the one in Go.

[What's A* algorithm](https://fr.wikipedia.org/wiki/Algorithme_A*)

[Visualize](https://www.angeluriot.com/maze_solver/)

[C++ repository](https://github.com/P0SlX/Barjokart)
### 🗺️ Prerequisites

You need golang 1.21 for running this project.

<p align="right"><a href="#top">⬆️</a></p>

<!-- USAGE EXAMPLES -->
## 💾 Usage

To be more user-friendly, a Makefile is present at the root of the project.
Even if you run with go run, you need to be at the root of the project and using img parameter with path of image you'd like to analyze.

### ⚙️ Using makefile
```shell
$ make run img=ressources/first_level.png # Run the application using go run under the hood
$ make run-build img=ressources/first_level.png # Run the application from dist folder, build it before !
$ make clean-cache # Clean the cache of the project
$ make build # Run the clean cache before building
$ make time # Time the execution of the application
$ make benchmark # Run the benchmark of the application, from benchmark folder
```

### 🏃 Using go run
```shell
$ go run main.go img=path_to_img # Run the application
```

### 📺Screenshots

Couleur selon H

<img src="https://i.imgur.com/j8gsEYY.png" alt="Logo" width="160" height="160">

Bidirectionnel

<img src="https://i.imgur.com/TmK49hq.png" alt="Logo" width="160" height="160">

<p align="right"><a href="#top">⬆️</a></p>

### 📊 Improvement
- `Node.Already` for flag already visited nodes | **-25%**
   - Avoid to iterate over a list of several hundred/thousand nodes at each iteration
- `PriorityQueue` make us gain an exponential time, we found that it make us gain MINIMUM 10 000% on execution time

### 🧙 What's next, Bidirectionnel ?

We would implement a bidirectional A* algorithm. The idea is to run two A* algorithms at the same time, one from the start and one from the end.
Unfortunately, we didn't have time to finish, and it's still bugged.
However, we have some results :
   - On very small images, the launch of A* in bidirectional sometimes has no performance gain compared to classic A*.
     - Indeed, depending on the system scheduler, it may be that a routine ends before the other has even started.
   - On larger images, we observe a gain of up to 70% in some cases.

He would be interesting to base ourselves on these [works](https://core.ac.uk/download/pdf/19184689.pdf) to finalize / improve the bidirectional A* algorithm.

<!-- CONTRIBUTING -->
## 🔗 Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right"><a href="#top">⬆️</a></p>


<!-- LICENSE -->
## 📰 License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right"><a href="#top">⬆️</a></p>


<!-- CONTACT -->
## 📫 Contact

Reach our at : gauron.dorian.pro@gmail.com, florian.savoure@gmail.com .

Project Link: [https://github.com/P0SlX/go-star](https://github.com/P0SlX/go-star)

<p align="right"><a href="#top">⬆️</a></p>


<!-- ACKNOWLEDGMENTS -->
## ⛱️ Acknowledgments

This space is a list to resources i found helpful and would like to give credit to.

* [![Go Badge](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
* [Bidirectionnal A*](https://core.ac.uk/download/pdf/19184689.pdf)
* [Pprof](https://github.com/google/pprof)
* [Benchmark](https://www.practical-go-lessons.com/chap-34-benchmarks)
* [Hyperfine](https://github.com/sharkdp/hyperfine)

<p align="right"><a href="#top">⬆️</a></p>

<a href="https://github.com/othneildrew/Best-README-Template">Template inspired by othneildrew</a>
