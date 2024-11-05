package aoc.y2018.d01; 

import org.junit.Test;
import static org.junit.Assert.*;

public class SolutionTest {
    @Test public void solutionHasAGreeting() {
        Solution classUnderTest = new Solution();
        assertNotNull("app should have a greeting", classUnderTest.getGreeting());
    }
}
