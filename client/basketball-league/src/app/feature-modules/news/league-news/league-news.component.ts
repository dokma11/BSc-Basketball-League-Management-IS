import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { User } from 'src/app/infrastructure/auth/model/user.model';
import { Trade } from 'src/app/shared/model/trade.model';
import { NewsService } from '../news.service';

@Component({
  selector: 'app-league-news',
  templateUrl: './league-news.component.html',
  styleUrls: ['./league-news.component.css'],
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
export class LeagueNewsComponent implements OnInit{
  user: User | undefined;
  backgroundSize: string = '100% 100%';
  trades: Trade[] = [];

  constructor(private newsService: NewsService,
              private snackBar: MatSnackBar,) { }

  ngOnInit(): void {
    this.getNews();
  }

  getNews() {
    this.newsService.getAllTrades().subscribe({
      next: (result: Trade[] | Trade) => {
        if(Array.isArray(result)) {
          this.trades = result;
        }
      }
    })
  }

  handleDialogClosed(result: any) {
    this.getNews();
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }
}
