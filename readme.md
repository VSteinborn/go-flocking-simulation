# go-flocking-simulation

A simple simulation of the [Boids flocking algorithm](https://en.wikipedia.org/wiki/Boids), meant to simulate the flocking behavior of birds in groups.

![Boids Simulation Gif](movie.gif)

The simulation is written in the Go programming language, and the results are plotted using Python via Matplotlib.

## Go Simulation

Run the Go simulation via the following command:

```shell
go run .
```

This will write the positions of the boids to the `./out/` directory.

### Performance

It takes roughly 7s to simulate 200 boids over 1000 time steps.

## Plotting

Plot the positions of the boids in an animated gif file by running the following python command:

```shell
python plot-trajectory.py
```

This will create a `movie.gif` file of the desired animated plot.
