import matplotlib.pyplot as plt
import matplotlib.patches as patches
import numpy as np
import os
import matplotlib.animation as animation


def load_positions():
    position_files = get_position_files()
    arrays = [np.genfromtxt(fname, delimiter=",") for fname in position_files]
    return np.stack(arrays, axis=0)


def get_position_files():
    out_dir = "./out"
    return [
        os.path.join(out_dir, file_name)
        for file_name in sorted(os.listdir(path=out_dir))
    ]


def init_plot(positions):
    fig, ax = plt.subplots()

    scat = ax.scatter(positions[:, 0], positions[:, 1], c="b", s=5)
    ax.set(xlim=[-0.25, 1.25], ylim=[-0.25, 1.25], xlabel="X [unit]", ylabel="Y [unit]")

    rect = patches.Rectangle((0, 0), 1, 1, linewidth=1, edgecolor="r", facecolor="none")
    ax.add_patch(rect)
    return fig, scat


def plot_boids():
    positions = load_positions()

    fig, scat = init_plot(positions[0])

    def update(frame):
        scat.set_offsets(positions[frame])
        return scat

    ani = animation.FuncAnimation(
        fig=fig, func=update, frames=positions.shape[0]#, interval=0.03
    )
    ani.save("movie.gif",fps=60)# gif is slow for some reason


if __name__ == "__main__":
    plot_boids()
