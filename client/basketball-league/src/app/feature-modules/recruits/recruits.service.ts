import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Employee } from 'src/app/shared/model/employee.model';
import { InterviewProposal } from 'src/app/shared/model/interviewProposal.model';
import { Recruit } from 'src/app/shared/model/recruit.model';
import { Team } from 'src/app/shared/model/team.model';
import { TrainingProposal } from 'src/app/shared/model/trainingProposal.model';
import { environment } from 'src/env/environment';

@Injectable({
  providedIn: 'root'
})
export class RecruitsService {

  constructor(private http: HttpClient) { }

  getAllRecruits(): Observable<Recruit> {
    return this.http.get<Recruit>(environment.host + 'recruit');
  }

  createInterviewProposal(proposal: InterviewProposal): Observable<InterviewProposal> {
    return this.http.post<InterviewProposal>(environment.host + 'interviewRequest', proposal);
  }

  getInterviewProposalsByRecruitId(recruitId: number): Observable<InterviewProposal> {
    return this.http.get<InterviewProposal>(environment.host + 'interviewRequest-receiver/' + recruitId);
  }

  createTrainingProposal(proposal: TrainingProposal): Observable<TrainingProposal> {
    return this.http.post<TrainingProposal>(environment.host + 'trainingRequest', proposal);
  }

  getTrainingProposalsByRecruitId(recruitId: number): Observable<TrainingProposal> {
    return this.http.get<TrainingProposal>(environment.host + 'trainingRequest-receiver/' + recruitId);
  }

  getTeamByCoachID(coachId: number): Observable<Team> {
    return this.http.get<Team>(environment.host + 'team-user/' + coachId);
  }

  getCoachByID(coachId: number): Observable<Employee> {
    return this.http.get<Employee>(environment.host + 'user/' + coachId);
  }

  updateInterviewProposal(proposal: InterviewProposal): Observable<InterviewProposal> {
    return this.http.put<InterviewProposal>(environment.host + 'interviewRequest', proposal);
  }

  updateTrainingProposal(proposal: TrainingProposal): Observable<TrainingProposal> {
    return this.http.put<TrainingProposal>(environment.host + 'trainingRequest', proposal);
  }

  declareForDraft(recruit: Recruit): Observable<Recruit> {
    return this.http.post<Recruit>(environment.host + 'recruit', recruit);
  }

  getAllRecruitsByName(name: string): Observable<Recruit> {
    return this.http.get<Recruit>(environment.host + 'recruit-name/' + name);
  }

  mapReverseSearch(lat: number, lon: number): Observable<any> {
    return this.http.get(
      `https://nominatim.openstreetmap.org/reverse?format=json&lat=${lat}&lon=${lon}&<params>`
    );
  }
}
