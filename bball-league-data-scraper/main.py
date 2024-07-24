import requests
import bs4

team_abbreviations = {
            'Brooklyn Nets': 'BRK',     # Switched from NJN to BRK
            'Golden State Warriors': 'GSW',
            'Los Angeles Lakers': 'LAL',
            'Los Angeles Clippers': 'LAC',
            'New Orleans Pelicans': 'NOP',
            'New York Knicks': 'NYK',
            'Oklahoma City Thunder': 'OKC',
            'San Antonio Spurs': 'SAS',
            'Boston Celtics': 'BOS',
            'Denver Nuggets': 'DEN',
            'Minnesota Timberwolves': 'MIN',
            'Cleveland Cavaliers': 'CLE',
            'Philadelphia 76ers': 'PHI',
            'Phoenix Suns': 'PHO',
            'Sacramento Kings': 'SAC',
            'Indiana Pacers': 'IND',
            'Dallas Mavericks': 'DAL',
            'Miami Heat': 'MIA',
            'Orlando Magic': 'ORL',
            'Chicago Bulls': 'CHI',
            'Atlanta Hawks': 'ATL',
            'Toronto Raptors': 'TOR',
            'Charlotte Hornets': 'CHO',     # Switched from CHA to CHO
            'Washington Wizards': 'WAS',
            'Detroit Pistons': 'DET',
            'Utah Jazz': 'UTA',
            'Houston Rockets': 'HOU',
            'Memphis Grizzlies': 'MEM',
            'Portland Trail Blazers': 'POR',
            'Milwaukee Bucks': 'MIL',
}

base_url = 'https://www.basketball-reference.com/'


def scrape_players():
    for team, abbreviation in team_abbreviations.items():
        response = requests.get(base_url + 'teams/' + abbreviation + '/2025.html#all_roster')
        response.raise_for_status()
        soup = bs4.BeautifulSoup(response.text, 'html.parser')

        table = soup.find('table', {'id': 'roster'})

        trs = table.find_all('tr')

        for tr in trs:
            tds = tr.find_all('td')
            if len(tds) > 0:
                # Split into first name and last name
                name_parts = tds[0].find('a').get_text().split(' ')
                first_name = name_parts[0]
                if len(name_parts) > 2:
                    last_name = ' '.join(name_parts[1:])
                else:
                    last_name = name_parts[1]

                # There are no precise position information, so I have to scrape positions manually from player's page
                a_tag = tds[0].find('a')
                player_page = base_url + a_tag.get('href')
                player_page_response = requests.get(player_page)
                player_page_response.raise_for_status()
                player_page_soup = bs4.BeautifulSoup(player_page_response.text, 'html.parser')

                position_div = player_page_soup.find('div', {'id': 'meta'})
                p_tags = position_div.find_all('p')
                for p in p_tags:
                    if p.get_text().__contains__('Position'):
                        if p.get_text().__contains__('Point Guard'):
                            print('Point Guard')
                            # TODO handle PG actions:

                        elif p.get_text().__contains__('Shooting Guard'):
                            print('Shooting Guard')
                            # TODO handle SG actions:

                        elif p.get_text().__contains__('Small Forward'):
                            print('Small Forward')
                            # TODO handle SF actions:

                        elif p.get_text().__contains__('Power Forward'):
                            print('Power Forward')
                            # TODO handle PF actions:

                        elif p.get_text().__contains__('Center'):
                            print('Center')
                            # TODO handle C actions:

                # Ovde moram da napravim poseban switch recimo (ili nesto mnogo slicno) pa da tako onda odredim string za to
                position = tds[1].get_text() # Ovaj ne daje dovoljno informacija o poziciji...

                # Convert height from feet and inches to centimeters
                height_parts = tds[2].get_text().split('-')     # Split the string for easier conversion
                height = int(int(height_parts[0]) * 30.48 + int(height_parts[1]) * 2.54)

                # Convert weight from pounds to kilograms
                if len(tds[3].get_text()) > 0:
                    weight = str(int(float(tds[3].get_text()) * 0.453592))
                else:
                    weight = 'Unknown'  # There is no information about the weight on the basketball-reference page

                # Ovde moram da pretvorim u neki format za datume mislim da ce mi to biti jasno tek kada krenem sa implementacijom bekenda
                birthday = tds[4].get_text()

                # These attributes don't value to scraping that much, so they'll be very simple
                password = last_name + '123'    # Dodati hesiranje....
                email = first_name + '@gmail.com'
                # TODO: Dodati inserte na kraju za korisnika, zaposlenog i igraca :)))))))))))


def scrape_teams():
    team_id = 0
    for team, abbreviation in team_abbreviations.items():
        response = requests.get(base_url + 'teams/' + abbreviation)
        response.raise_for_status()
        soup = bs4.BeautifulSoup(response.text, 'html.parser')

        div_parent_tag = soup.find_all('div', {'id': 'meta'})
        for div in div_parent_tag:
            span_tags = div.find_all('span')
            for span in span_tags:
                if len(span.get_text()) > 0:
                    team_name = span.get_text()
                    print(team_name)    # Radi provere

            p_tags = div.find_all('p')
            for p in p_tags:
                if p.get_text().__contains__('Location'):
                    team_location = p.get_text()
                    print(team_location)    # Radi provere
                if p.get_text().__contains__('Seasons'):
                    # Get the establishment year from the string
                    p_parts = p.get_text().split(';')
                    p_parts[1] = p_parts[1].strip().split()
                    team_establishment_year = p_parts[0].split('-')[0]
                    print(team_establishment_year)    # Radi provere

        print(team_id)    # Radi provere
        # TODO: Razmotriti da li ovde pisati inserte za bazu
        team_id = team_id + 1


def scrape_coaches():
    for team, abbreviation in team_abbreviations.items():
        response = requests.get(base_url + 'teams/' + abbreviation)
        response.raise_for_status()
        soup = bs4.BeautifulSoup(response.text, 'html.parser')

        table = soup.find('table', _class='sortable stats_table now_sortable sticky_table eql rel lel')
        tr_tags = table.find_all('tr')
        a_tags = tr_tags[0].find_all('a')
        if a_tags:
            # Manually go to the coaches page for more information
            coach_page = base_url + 'coaches/' + a_tags[-1].get('href')
            coach_page_response = requests.get(coach_page)
            coach_page_response.raise_for_status()
            coach_page_soup = bs4.BeautifulSoup(coach_page_response.text, 'html.parser')

            span_tags = coach_page_soup.find_all('span')
            coach_name_parts = span_tags[-1].get_text().split(' ')
            coach_first_name = coach_name_parts[0]
            coach_last_name = coach_name_parts[1]

            birthday_span_tag = coach_page_soup.find('span', {'id': 'necro-birth'})
            coach_birthday = birthday_span_tag.get_text()

            table = coach_page_soup.find('table', _class='suppress_glossary sortable stats_table now_sortable')
            th_tags = table.find_all('th')
            coach_years_of_experience = 0
            for th in th_tags:
                if len(th.get_text()) > 0:
                    th_parts = th.get_text().split(' ')
                    coach_years_of_experience += int(th_parts[0])


if __name__ == "__main__":
    print('Scraping players... \n')
    scrape_players()
    print('Player scraping finished!\n')
    print('Scraping teams... \n')
    scrape_teams()
    print('Team scraping finished\n')
    print('Scraping coaches... \n')
    scrape_coaches()
    print('Coach scraping finished\n')
