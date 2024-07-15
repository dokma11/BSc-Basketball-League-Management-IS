import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { AuthService } from 'src/app/infrastructure/auth/auth.service';
import { User } from 'src/app/infrastructure/auth/model/user.model';

@Component({
  selector: 'app-league-news',
  templateUrl: './league-news.component.html',
  styleUrls: ['./league-news.component.css']
})
export class LeagueNewsComponent implements OnInit{
  user: User | undefined;
  backgroundSize: string = '100% 100%';
  //requests: PersonalTourRequest[] = [];   ovde treba da budu novosti

  private dialogRef: any;

  constructor(private authService: AuthService,
              private dialog: MatDialog,
              private snackBar: MatSnackBar,) {

  }

  ngOnInit(): void {
    this.getNews();
  }

  getNews() {
    // TODO: Get all the trade news from be 
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
