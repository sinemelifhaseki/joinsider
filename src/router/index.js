import Leagues from "../components/Leagues";
import LeagueDetails from "../components/LeagueDetails";
import TeamDetail from "../components/TeamDetail";

export default [
  {path: '/', component: TeamDetail},
  {path: '/leagues/:league_id', component: LeagueDetails},
  {path: '/leagues/:league_id/teams/:team_id', component: TeamDetail},
]
