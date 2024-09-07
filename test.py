from tkinter import Tk, Canvas
from PIL import Image, ImageTk

class CoordinateTool:
    def __init__(self, image_path):
        self.root = Tk()
        self.root.title("Coordinate Tool")

        self.image = Image.open(image_path)
        self.tk_image = ImageTk.PhotoImage(self.image)

        self.canvas = Canvas(self.root, width=self.tk_image.width(), height=self.tk_image.height())
        self.canvas.pack()

        self.canvas.create_image(0, 0, anchor='nw', image=self.tk_image)
        self.canvas.bind("<Button-1>", self.get_coordinates)

        self.root.mainloop()

    def get_coordinates(self, event):
        print(f"Coordinates: ({event.x}, {event.y})")

if __name__ == "__main__":
    CoordinateTool('screeenshot.png')
