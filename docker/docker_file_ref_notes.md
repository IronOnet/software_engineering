# Dockerfile

Docker runs instructions in a Dockerfile in order. A Dockerfile must begin
with a FROM instruction. This may be after parser directives, comments and
globally scoped ARGS.

The FROM instruction specifies the Parent image from which you are building.
FROM may only be preceded by one or more ARG instructions, which declare
arguments that are used in FROM lines in the Dockerfile.

Docker treats lines that begin with # as a comment, unless the line is a valid
parser directive. A # marker anywhere else in a line is treated as an argument.
