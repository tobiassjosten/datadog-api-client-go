import yaml
import sys


def main():
    f = sys.argv[1]
    with open(f) as o:
        data = yaml.load(o)

    data["version"] = 2

    for i, interaction in enumerate(data["interactions"]):
        interaction["id"] = i
        interaction["response"]["duration"] = "100ms"

    output = yaml.dump(data)
    with open(f, "w") as o:
        o.write(output)


if __name__ == "__main__":
    main()
