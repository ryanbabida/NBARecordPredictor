create database predictor

create table [if not exists] records {
    games_played integer
    wins integer
    losses integer
    win_percentage float
    points float
    field_goals_made float
    field_goals_attempted float
    field_goal_percentage float
    threes_made float
    threes_attempted float
    three_percentage float
    free_throws_made float
    free_throws_attempted float
    free_throw_percentage float
    offensive_rebounds float
    defensive_rebounds float
    rebounds float
    assists float
    turnover float
    steals float
    blocks float
    blocks_against float
    personal_fouls float
    personal_fouls_against float
    plus_minus float
}