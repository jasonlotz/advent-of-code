package aoc.y2018.d01;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class Solution {
  // Run from the root of the root of the Gradle project
  private static Path basePath = Paths.get(System.getProperty("user.dir"));
  private static final Path INPUT_FILE = basePath.resolve("../../input-files/2018/01/input.txt");

  private List<String> readFile() {
    try {
      return Files.readAllLines(INPUT_FILE);
    } catch (IOException e) {
      throw new RuntimeException("Failed to read file");
    }
  }

  public int sum(List<String> lines) {
    return lines.stream().mapToInt(Integer::parseInt).sum();
  }

  public int findDuplicateFreq(List<String> lines) {
    int totalFreq = 0;
    int seenTwice = 0;
    Set<Integer> seenFreqs = new HashSet<>();

    // Loop through the list of frequencies until we find a duplicate
    // frequency even if it means looping through the list multiple times
    while (true) {
      for (String line : lines) {
        int freq = Integer.parseInt(line);

        totalFreq += freq;

        if (seenFreqs.contains(totalFreq)) {
          seenTwice = totalFreq;
          break;
        }

        seenFreqs.add(totalFreq);
      }

      if (seenTwice != 0) {
        break;
      }
    }

    return seenTwice;
  }

  public static void main(String[] args) {
    Solution solution = new Solution();

    List<String> lines = solution.readFile();

    int totalFreq = solution.sum(lines);
    int seenTwice = solution.findDuplicateFreq(lines);

    System.out.println("Part 1: " + totalFreq);
    System.out.println("Part 2: " + seenTwice);
  }
}
