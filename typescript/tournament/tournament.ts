export class Tournament {
  private teams: Record<string, Team> = {};
  private reportBuilder: TournamentReportBuilder | undefined;
  constructor() {
    this.reportBuilder = new TournamentReportBuilder(this);
  }

  public tally(input: string): string {
    this.insertMatches(input);

    return this.getReportBuilder().build();
  }

  private getReportBuilder(): TournamentReportBuilder {
    if (!this.reportBuilder) {
      throw new Error("Report builder not initialized");
    }

    return this.reportBuilder;
  }

  private insertMatches(input: string): void {
    if (!input.trim()) {
      return;
    }

    input.split("\n").forEach((line) => {
      const [team1, team2, result] = line.split(";");

      if (!isResult(result)) {
        throw new Error("Invalid result");
      }

      const match = new Match(team1, team2, result);
      this.getTeamByName(team1).addMatchPlayed(match);
      this.getTeamByName(team2).addMatchPlayed(match);
    });
  }

  private getTeamByName(name: string): Team {
    if (!this.teams[name]) {
      this.teams[name] = new Team(name);
    }

    return this.teams[name];
  }

  public getTeamScore(team: Team): number {
    return (
      team.getWins() * PointsEnum.WIN +
      team.getDraws() * PointsEnum.DRAW +
      team.getLosses() * PointsEnum.LOSS
    );
  }

  public getSortedTeams(): Team[] {
    const teams = Object.values(this.teams);

    return teams.sort((a, b) => {
      if (this.getTeamScore(a) !== this.getTeamScore(b)) {
        return this.getTeamScore(b) - this.getTeamScore(a);
      }

      return a.getName().localeCompare(b.getName());
    });
  }
}

type CellValue = string | number;
type ReportLine = [
  CellValue,
  CellValue,
  CellValue,
  CellValue,
  CellValue,
  CellValue
];

class TournamentReportBuilder {
  constructor(private tournament: Tournament) {}

  private buildHeader(): string {
    return this.createReportLine("Team", "MP", "W", "D", "L", "P");
  }

  private buildRow(team: Team): string {
    return this.createReportLine(
      team.getName(),
      team.getMatchesPlayed(),
      team.getWins(),
      team.getDraws(),
      team.getLosses(),
      this.tournament.getTeamScore(team)
    );
  }

  private createReportLine(...reportLine: ReportLine): string {
    return reportLine
      .map((cellValue, index, self) => {
        if (index === 0 && typeof cellValue === "string") {
          return this.getCellTeamName(cellValue);
        }

        if (index === self.length - 1) {
          return this.getCell(cellValue).trimEnd();
        }

        return this.getCell(cellValue);
      })
      .join("|");
  }

  private getCell(value: CellValue): string {
    return value.toString().padStart(3, " ") + " ";
  }

  private getCellTeamName(value: string): string {
    return value.padEnd(31, " ");
  }

  public build(): string {
    const header = this.buildHeader();
    const sortedTeams = this.tournament.getSortedTeams();
    const teamLines = sortedTeams.map((team) => this.buildRow(team));

    return [header, teamLines.join("\n")].filter(Boolean).join("\n");
  }
}

class Team {
  private matchesPlayed: number = 0;
  private wins: number = 0;
  private draws: number = 0;
  private losses: number = 0;

  constructor(private name: string) {}

  public addMatchPlayed(match: Match): void {
    this.matchesPlayed++;
    const result = match.getTeamResult(this);

    if (result === ResultEnum.WIN) {
      this.wins++;
    } else if (result === ResultEnum.DRAW) {
      this.draws++;
    } else {
      this.losses++;
    }
  }

  public getName(): string {
    return this.name;
  }

  public getMatchesPlayed(): number {
    return this.matchesPlayed;
  }

  public getWins(): number {
    return this.wins;
  }

  public getDraws(): number {
    return this.draws;
  }

  public getLosses(): number {
    return this.losses;
  }
}

class Match {
  constructor(
    public team1: string,
    public team2: string,
    public result: ResultEnum
  ) {}

  public getTeamResult(team: Team): ResultEnum {
    if (team.getName() == this.team1) {
      return this.result;
    }

    if (team.getName() == this.team2) {
      if (this.result === ResultEnum.WIN) {
        return ResultEnum.LOSS;
      }

      if (this.result === ResultEnum.LOSS) {
        return ResultEnum.WIN;
      }

      return ResultEnum.DRAW;
    }

    throw new Error("Invalid team");
  }
}

enum ResultEnum {
  WIN = "win",
  DRAW = "draw",
  LOSS = "loss",
}

function isResult(v: any): v is ResultEnum {
  return (
    typeof v === "string" && Object.values(ResultEnum).some((val) => val === v)
  );
}

enum PointsEnum {
  WIN = 3,
  DRAW = 1,
  LOSS = 0,
}
