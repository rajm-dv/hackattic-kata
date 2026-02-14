def countDay(n):
    days = ["Thursday", "Friday", "Saturday", "Sunday", "Monday", "Thesday", "Wednesday"]
    print(days[((n % 7) + 7) % 7])


def main():
    while 1:
        try:
            line = int(input())
            countDay(line)
        except EOFError:
            break
