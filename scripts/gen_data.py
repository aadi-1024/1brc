import random
import string
import sys

OUT_PATH = None
RES_PATH = None

def main():
    if sys.argv.__len__() < 4:
        print('not all outputs given: [output] [result] [n]')
        return
    
    OUT_PATH = sys.argv[1]
    RES_PATH = sys.argv[2]
    try:
        n = int(sys.argv[3])
    except:
        n = 1_000_000_000
    print(f"generating file with {n} rows")

    #generate 10k 8 letter names along with their mean and var to use for a gaussian distribution
    names = set()
    mean = dict()
    var = dict()

    while len(names) != 10000:
        name = ''.join(random.choices(string.ascii_lowercase, k=10))
        if names.__contains__(name):
            continue
        names.add(name)
        mean[name] = random.uniform(0, 100)
        var[name] = random.uniform(10, 30)

    names = list(names)

    #store final results
    results: dict[str, dict[str, float]] = {}

    #given that random.choice follows a uniform distribution
    #all of the names would appear roughly the same number of
    #times in the final output
    with open(OUT_PATH, 'w') as file:
        for i in range(n):
            name = random.choice(names)
            val = random.gauss(mean[name], var[name])

            file.write(f'{name},{val:.2f}\n')
            if results.__contains__(name):
                results[name]['total'] += val
                results[name]['n'] += 1
            else:
                results[name] = {
                    'total': val,
                    'n': 1,
                    'mean': 0.0
                }
    
    with open(RES_PATH, 'w') as file:
        for k, v in results.items():
            v['mean'] = v['total'] / v['n']
            file.write(f'{k},{v['total']},{v['mean']:.2f}\n')

if __name__ == "__main__":
    main()
