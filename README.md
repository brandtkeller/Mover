# Mover
Companion application for your Autonomous sit-stand desk


## Purpose
I own a sit-stand desk [Autonomous]() and as much as the features it has are great for my body - deep focus can often lead to a general lack of awareness for extended periods of time where I don't utilize its features.

Let's see if we can't fix that by adding a middle-man between the buttons that I push to change the positions and the motor controllers to start scheduling a proper balance of sit/stand time without my intervention. 

## Initial Goal
- Connect to system without having to splice any autonomous wiring
- Raspberry Pi w/ 2x relays wired
- API for exposing Start/Stop functionality
- Randomized periods of sitting/standing that is aligned with general health recommendations

## Future Goals
- Synchronize with Calendar to be able to tell when I might be out-of-office and it doesn't have to move
    - Fully automated start/stop
    - Still have override functionality above
    - Maybe enable some checking to only change if not in a meeting?