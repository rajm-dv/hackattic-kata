def compressed(input):
    n = len(input)
    i = 0
    while i < n:
        cnt = 1
        curr = input[i]
        while i + cnt < n and input[i + cnt] == curr:
            cnt += 1
        if cnt > 2:
            print(f"{cnt}{curr}", end="")
        else:
            for j in range(cnt):
                print(curr, end="")
        i += cnt


def main():
    while 1:
        try:
            line = input()
            compressed(line)
            print("")
        except EOFError:
            break


if __name__ == "__main__":
    main()
