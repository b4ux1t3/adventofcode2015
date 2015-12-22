import java.nio.file.Files;
import java.nio.file.Paths;
import java.nio.file.Path;
import java.io.*;

public class notlisp
{
	public static void main(String[] args)
	{
		// I'm thinking we need to establish that Santa is on the ground floor.
		int santaPos = 0;

		// Chris: Read the instructions from "input"
		Path path = Paths.get("input");

		byte[] input;
		try
		{
			input = Files.readAllBytes(path);
		} catch(IOException err)
		{
			input = null;
		}

		// So it needs to count how many of each kind. [of instruction]
		// We need to declare them....and set them equal to something.
		int up = 0;
		int down = 0;
		for (int i = 0; i < input.length; i++)
		{
			if (input[i] == 40) // '(' is 40 in ASCII
			{
				up++;
			}
			else
			{
				down++;
			}
		}
		// I feel like its too obvious so it can't be right. Do we find the difference??
		// And that would be the floor he needs to go to.
		int result = up - down;
		System.out.printf("%d\n", result);
	}
}
