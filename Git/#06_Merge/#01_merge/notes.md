# MERGE

"What's the point of having mulitple branches?" you might ask. They're most often used to safely make changes without affecting your(or your _team's_) primary branch. However, once you're happy with your changes, you'll wnat ot merge them back into the main branch so that they make their way into the final product.

### VISUAL

Let's say you're in a state where you have two branches, each with their own unique commits:

    A - B - C   main
     \
      D - E     other_brach

If you merge `other_brach` into `main`, Git combines both branches by creating a new commit that has _both_ histories as parents. In the diagram below, `F` is a merge commit that has `C` and `E` as parents. `F` brings all the changes from `D` and `E` back into the `main` branch.

    A - B - C - F   main
     \      /
      D - E         other_brach
