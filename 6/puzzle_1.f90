PROGRAM read_data
    IMPLICIT NONE

    INTEGER, PARAMETER :: n = 4  ! Number of data points
    INTEGER :: i, io_status, product = 1
    INTEGER :: time_values(n), distance_values(n)
    CHARACTER(len=10) :: time_label, distance_label
    REAL :: t_charge_min, t_charge_max

    ! Open the file
    OPEN(UNIT=10, FILE='puzzle_input.txt', STATUS='OLD', ACTION='READ', IOSTAT=io_status)
    IF (io_status /= 0) THEN
        PRINT *, 'Error opening file!'
        STOP
    END IF

    ! Read the time values
    READ(10, *) time_label, (time_values(i), i=1,n)

    ! Read the distance values
    READ(10, *) distance_label, (distance_values(i), i=1,n)

    CLOSE(10)

    DO i = 1, n
        t_charge_min = time_values(i) / 2.0 - SQRT(-4.0 * distance_values(i) + time_values(i) ** 2) / 2.0 + 0.00001
        t_charge_max = time_values(i) / 2.0 + SQRT(-4.0 * distance_values(i) + time_values(i) ** 2) / 2.0 - 0.00001

        t_charge_min = CEILING(t_charge_min)
        t_charge_max = FLOOR(t_charge_max)

        product = product * INT(t_charge_max - t_charge_min + 1)
    END DO

    PRINT *, 'Day 6 Puzzle 1', product

END PROGRAM read_data
