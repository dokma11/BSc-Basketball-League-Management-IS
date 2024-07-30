import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, Input, OnInit } from '@angular/core';
import { faCheck, faTimes, faPen, faTrash } from '@fortawesome/free-solid-svg-icons';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { Trade } from 'src/app/shared/model/trade.model';
import { NewsService } from '../news.service';
import { TradeProposal } from 'src/app/shared/model/tradeProposal.model';
import { TradesService } from '../../trades/trades.service';
import { Team } from 'src/app/shared/model/team.model';

@Component({
  selector: 'app-news-card',
  templateUrl: './news-card.component.html',
  styleUrls: ['./news-card.component.css'],
  animations: [
      trigger("fadeIn", [
        transition(":enter", [
            style({ opacity: 0, transform: "translateX(-40px)" }),
            animate(
                "0.5s ease",
                style({ opacity: 1, transform: "translateX(0)" }),
            ),
        ]),
      ]),
      trigger('buttonState', [
        state('clicked', style({
          transform: 'scale(0.9)',
          opacity: 0.5
        })),
        transition('* => clicked', [
          animate('200ms')
        ]),
        transition('clicked => idle', [
          animate('200ms')
        ])
      ]),
  ],
})
export class NewsCardComponent implements OnInit{
  headline: string = ''; 
  @Input() trade!: Trade;
  user: User | undefined;
  tradeOccurrenceDate: string = '';
  tradeOccurrenceTime: string = '';
  journalist: string = '';
  dateDay: string = '';
  dateMonth: string = '';
  dateYear: string = '';
  description: string = '';
  firstTeam: Team | undefined;
  secondTeam: Team | undefined;
  firstTeamBgColor: string = 'olivedrab';
  secondTeamBgColor: string = 'red';
  imageSource: string = 'https://imgs.search.brave.com/udmDGOGRJTYO6lmJ0ADA03YoW4CdO6jPKGzXWvx1XRI/rs:fit:860:0:0/g:ce/aHR0cHM6Ly90My5m/dGNkbi5uZXQvanBn/LzAyLzY4LzU1LzYw/LzM2MF9GXzI2ODU1/NjAxMl9jMVdCYUtG/TjVyalJ4UjJleVYz/M3puSzRxblllS1pq/bS5qcGc';
  tradeProposal: TradeProposal | undefined;

  teamColors: { [key: string]: string } = {
    'Atlanta Hawks': '#E03A3E',
    'Boston Celtics': '#007A33',
    'Brooklyn Nets': '#000000',
    'Charlotte Hornets': '#1D1160',
    'Chicago Bulls': '#CE1141',
    'Cleveland Cavaliers': '#860038',
    'Dallas Mavericks': '#00538C',
    'Denver Nuggets': '#0E2240',
    'Detroit Pistons': '#C8102E',
    'Golden State Warriors': '#1D428A',
    'Houston Rockets': '#CE1141',
    'Indiana Pacers': '#002D62',
    'Los Angeles Clippers': '#C8102E',
    'Los Angeles Lakers': '#552583',
    'Memphis Grizzlies': '#5D76A9',
    'Miami Heat': '#98002E',
    'Milwaukee Bucks': '#00471B',
    'Minnesota Timberwolves': '#0C2340',
    'New Orleans Pelicans': '#0C2340',
    'New York Knicks': '#006BB6',
    'Oklahoma City Thunder': '#007AC1',
    'Orlando Magic': '#0077C0',
    'Philadelphia 76ers': '#006BB6',
    'Phoenix Suns': '#1D1160',
    'Portland Trail Blazers': '#E03A3E',
    'Sacramento Kings': '#5A2D81',
    'San Antonio Spurs': '#C4CED4',
    'Toronto Raptors': '#CE1141',
    'Utah Jazz': '#002B5C',
    'Washington Wizards': '#002B5C'
  };

  constructor(private newsService: NewsService, 
              private tradesService: TradesService) { }

  ngOnInit(): void {
    this.newsService.getTradeProposalByID(this.trade.idZahTrg!).subscribe({
      next: (result: TradeProposal) => {
        this.tradeProposal = result;

        this.tradesService.getTeamByManagerID(this.tradeProposal.idMenadzerPos).subscribe({
          next: (result: Team) => {
            this.firstTeam = result;
          }
        });

        this.tradesService.getTeamByManagerID(this.tradeProposal.idMenadzerPrim!).subscribe({
          next: (result: Team) => {
            this.secondTeam = result;
          }
        });

        [this.tradeOccurrenceDate, this.tradeOccurrenceTime] = this.trade.datTrg.toString().split('T');
        [this.dateYear, this.dateMonth, this.dateDay] = this.tradeOccurrenceDate.split('-');
        this.tradeOccurrenceDate = this.dateDay + '.' + this.dateMonth + '.' + this.dateYear + '.'
    
        this.journalist = this.getRandomNbaInsider();
    
        if(this.journalist == 'Adrian Wojnarowski @wojespn') {
          this.imageSource = 'https://library.sportingnews.com/styles/twitter_card_120x120/s3/2023-06/Adrian%20Wojnarowski%20061923.jpg?itok=e5hXJqwV'
        } else if(this.journalist == 'Shams Charania @ShamsCharania') {
          this.imageSource = 'https://pbs.twimg.com/profile_images/1648391157045673984/PeoZeyFY_400x400.jpg'
        } else if(this.journalist == 'Rachel Nichols @Rachel_Nichols') {
          this.imageSource = 'https://media.npr.org/assets/img/2021/08/26/gettyimages-1234437863-ad4a1e7e1070c68a13d70d2bb0c12ae451ee5507.jpg'
        } else if(this.journalist == 'Brian Windhorst @WindhorstESPN') {
          this.imageSource = 'https://www.billboard.com/wp-content/uploads/2024/06/Brian-Windhorst-2023-billboard-1548.jpg?w=942&h=623&crop=1'
        } else if(this.journalist == 'Marc Stein @TheSteinLine') {
          this.imageSource = 'https://espnpressroom.com/us/files/2016/10/FOR-AMINA.jpg'
        } else if(this.journalist == 'Zach Lowe @ZachLowe_NBA') {
          this.imageSource = 'https://cdn.vox-cdn.com/thumbor/MfkWL090-RcOAJsiVQgeNlUZ650=/0x0:4368x2912/1200x800/filters:focal(1716x298:2414x996)/cdn.vox-cdn.com/uploads/chorus_image/image/71440573/1241592514.0.jpg'
        } else if(this.journalist == 'Chris Haynes @chrisbhaynes') {
          this.imageSource = 'https://pbs.twimg.com/profile_images/1760487569090097152/Bg04W_yB_400x400.jpg'
        }
    
        setTimeout(() => {
          if(this.trade.tipTrg == 0){ // PLAYER-PLAYER type
            this.description = this.getRandomPlayerPlayerDescription();
          } else if(this.trade.tipTrg == 1) { // PLAYER-PICK type
            this.description = this.getRandomPlayerPickDescription();
          } else if(this.trade.tipTrg == 2) { // PICK-PICK type
            this.description = this.getRandomPickPickDescription();
          }
      
          this.headline = this.getRandomHeadline();  

           this.firstTeamBgColor = this.teamColors[this.firstTeam?.nazTim!] || 'defaultColor';
           this.secondTeamBgColor = this.teamColors[this.secondTeam?.nazTim!] || 'defaultColor';

        }, 10);
      }
    });
  }

  getRandomNbaInsider(): string {
    const insiders = ['Adrian Wojnarowski @wojespn', 
      'Shams Charania @ShamsCharania', 
      'Rachel Nichols @Rachel_Nichols', 
      'Brian Windhorst @WindhorstESPN',
      'Marc Stein @TheSteinLine',
      'Zach Lowe @ZachLowe_NBA',
      'Chris Haynes @chrisbhaynes'
    ];
    const randomIndex = Math.floor(Math.random() * insiders.length);
    return insiders[randomIndex];
  }

  getRandomPlayerPlayerDescription(): string {
    const descriptions = [this.firstTeam?.nazTim + ' have traded a number of key players to the ' + this.secondTeam?.nazTim + ' and, in return, have received minimal compensation. Many analysts suggesting it could be a potentially significant misstep for them', 
      this.firstTeam?.nazTim + ' have traded a trio of bench players to the' + this.secondTeam?.nazTim + ' for a star forward. This trade is being hailed as a major victory for the Knicks, significantly strengthening their lineup and enhancing their chances for a deep playoff run.',
      this.firstTeam?.nazTim + ' have traded their star player for a plethora of young players from ' + this.secondTeam?.nazTim + ', signifying a potential rebuild. The fans are gonna love this one',
      this.firstTeam?.nazTim + ' have sent some of their top players to ' + this.secondTeam?.nazTim + ' and, in return, have received a promising rookie and significant future cap space', 
    ];
    const randomIndex = Math.floor(Math.random() * descriptions.length);
    return descriptions[randomIndex];
  }

  getRandomPlayerPickDescription(): string {
    const insiders = [this.firstTeam?.nazTim + ' have offloaded a number of their players to the ' + this.secondTeam?.nazTim + ', receiving only a few draft picks in return. Critics are strongly questioning the wisdom of this move', 
      this.firstTeam?.nazTim + ' have exchanged a number of their draft picks with the ' + this.secondTeam?.nazTim + ', receiving only a few role players in return. This could be a key move to improve their championship hopes',
      this.firstTeam?.nazTim + ' have offloaded a number of their future draft picks for a star-level player from the ' + this.secondTeam?.nazTim + '. This move has raised many eyebrows around the fanbase',
      this.firstTeam?.nazTim + ' have traded some of their current role players for a plethora of future draft picks from the ' + this.secondTeam?.nazTim + '. The fans can be happy because this is a great move for their future',
    ];
    const randomIndex = Math.floor(Math.random() * insiders.length);
    return insiders[randomIndex];
  }

  getRandomPickPickDescription(): string {
    const insiders = [this.firstTeam?.nazTim + ' have traded their first-round pick to the ' + this.secondTeam?.nazTim + ' in exchange for two future second-round picks. An under the radar move that could prove to be significant in the future', 
      this.firstTeam?.nazTim + ' have acquired a lottery pick from the ' + this.secondTeam?.nazTim + ' in exchange for their two late first-round picks. Many analysts are praising this move',
      this.firstTeam?.nazTim + ' have traded their first-round pick to the ' + this.secondTeam?.nazTim + ' in exchange for two future first-round picks. This could be a pragmatic approach to building depth for the future.',     
      this.firstTeam?.nazTim + ' have traded multiple second-round picks to the ' + this.secondTeam?.nazTim + ' for a future first-round pick. Some critics argue that the Rockets may have overpaid, but some seem to love it',
    ];
    const randomIndex = Math.floor(Math.random() * insiders.length);
    return insiders[randomIndex];
  }

  getRandomHeadline(): string {
    const insiders = [this.firstTeam?.nazTim + ' have partnered with the ' + this.secondTeam?.nazTim + ' for a trade', 
      this.firstTeam?.nazTim + ' finnally made a move!',
      this.firstTeam?.nazTim + ' have done a miracle.',     
      'What a move by the ' + this.firstTeam?.nazTim + '!',
      'What were the ' + this.firstTeam?.nazTim + ' thinking?',
      'An even trade between the ' + this.firstTeam?.nazTim + ' and the ' + this.secondTeam?.nazTim,
      'A surprising move by the ' + this.firstTeam?.nazTim,
      'A championship or bust move by the ' + this.firstTeam?.nazTim,
    ];
    const randomIndex = Math.floor(Math.random() * insiders.length);
    return insiders[randomIndex];
  }

  faCheck = faCheck;
  faTimes = faTimes;
  faPen = faPen;
  faTrash = faTrash;
}
