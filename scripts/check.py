import sys

def main():
    if len(sys.argv) < 3:
        print("need [expected] [output]")
        return

    expected = sys.argv[1]
    out = sys.argv[2]
    actual: dict[str, list[float]] = dict()

    #first read the actual results
    with open(expected, 'r') as og:
        for line in og.readlines():
            l = line.split(',')
            name = l[0]

            actual[name] = l[1:]

    with open(out, 'r') as out:
        for line in out.readlines():
            l = line.split(',')
            name = l[0]
            sum = float(l[1])
            mean = float(l[2])

            if abs(float(actual[name][0]) - sum) > 0.1 or abs(float(actual[name][1]) - mean) > 0.1:
                print(f'expected: {actual[name][0]},{actual[name][1]}, got {sum},{mean} for {name}')
                return False

    return True

if __name__ == "__main__":
    if main():
        print("check passed")
    else:
        print("check failed")
