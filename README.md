# uArm
Stuff for the uArm Swift Pro

G0 X250 Y0 Z130 F10000
G0 X1 Y1 Z130 F10000
G0 X0 Y0 Z100 F10000
G0 X100 Y100 Z100 F10000


// Get Current Mode, 0: normal, 1: laser, 2: 3D printing, 3: Universal Holder
P2400
// Set Current mode, mode index same as above
M2400

// Go to absolute location, XYZ, F = mm/minute
// G0 = laser off, G1 = laser on
// E22 = out of range
// X is towards or away from the central axis point, The minimun seems to change based on Y and Z
// Y is left and right. Looking at the uArm logo as the front, positive is towards the right, negative to the left
// Z starts at the plane of the base but seems likely to have some offset
G0 X100 Y100 Z100 F10000
G0 X150 Y0 Z100 F10000
G0 X200 Y-50 Z100 F10000
G0 X200 Y-50 Z15 F10000

// Check if coordinates are reachable
// P 0=Cartesian, 1=Polar
// Response 1=reachable, 0=unreachable
M2222 X100 Y100 Z100 P0


// Attach all the joint motors
// When a new connection is opened it seems like all joint motors are detached. This means the arm drops to the table, which is not the best.
M17

// Detach all motors
M2019

// Set current position as reference position
M2401

// Set the height zero point
M2410

// Set the offset of end-effector in mm
M2411 S100

## Queries

P<num>

P2200 = current angle of joints, output values are floats
ok B50 L50 R50

P2201 = device name
ok SwiftPro

P2202 = hardware version
ok V3.3.0

P2203 = software version
ok V4.9.0

P2204 = api version
ok V4.0.5

P2400 = check current status (0:normal, 1:laser, 2:3d 4:universal)
ok V1

// Timed Feedback
// Issue `M2120 V0.2` V is time in seconds, Will receive `@3` with values
M2120 V5
@3 X149.7157 Y0.0000 Z99.8552 R89.1000

// Use V0 to disable
M2120 V0


