from __future__ import print_function

outfile = None
last_date = None

with open("WoWCombatLog.txt", "r") as logfile:
    for row in logfile:
        date = row.split(' ')[0]

        if date != last_date:
            print("New date: {}".format(date))
            if outfile:
                outfile.close()

            outfile = open("wow_log__{}.txt".format(date.replace('/', '_')), "w")
            last_date = date


        outfile.write(row)

